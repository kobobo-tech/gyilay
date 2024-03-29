use crate::schematic::traits::{util::*, TraitImplementation};
use crate::workload_type::{ParamMap, REPLICATED_SERVICE_NAME, REPLICATED_TASK_NAME};
use k8s_openapi::api::{apps::v1 as apps, batch::v1 as batch};

use kube::client::APIClient;

/// A manual scaler provides a way to manually scale replicable objects.
#[derive(Clone, Debug)]
pub struct ManualScaler {
    pub name: String,
    pub instance_name: String,
    pub component_name: String,
    pub owner_ref: OwnerRefs,
    pub replica_count: i32,
    pub workload_type: String,
}

impl ManualScaler {
    pub fn from_params(
        name: String,
        inst_name: String,
        comp_name: String,
        params: ParamMap,
        owner_ref: OwnerRefs,
        workload_type: String,
    ) -> ManualScaler {
        info!("params: {:?}", &params);
        ManualScaler {
            name: name,
            instance_name: inst_name,
            component_name: comp_name,
            owner_ref: owner_ref,
            replica_count: params
                .get("replicaCount".into())
                .and_then(|p| p.as_i64().and_then(|i64| Some(i64 as i32)))
                .unwrap_or(1),
            workload_type: workload_type,
        }
    }

    fn scale(&self, ns: &str, client: APIClient) -> TraitResult {
        // TODO: We probably need to watch for the deployment to be created. Or this might be unnecessary.
        std::thread::sleep(std::time::Duration::from_secs(5));

        info!("Scaling {} to {:?}", &self.name, &self.replica_count);
        // It should be a safe assumption that we can look up every job and every
        // deployment with a particular Kubernetess name and update them appropriately.
        match self.workload_type.as_str() {
            REPLICATED_SERVICE_NAME => {
                let (req, _) = apps::Deployment::read_namespaced_deployment(
                    self.instance_name.as_str(),
                    ns,
                    Default::default(),
                )?;
                let res = client.request(req);
                if res.is_ok() {
                    let original: apps::Deployment = res.unwrap();
                    let dep = self.scale_deployment(original);

                    let (req2, _) = apps::Deployment::replace_namespaced_deployment(
                        self.instance_name.as_str(),
                        ns,
                        &dep,
                        Default::default(),
                    )?;
                    let res2: Result<serde_json::Value, failure::Error> = client.request(req2);
                    if res2.is_err() {
                        let err = res2.unwrap_err();
                        error!(
                            "Scaling error: {}",
                            serde_json::to_string_pretty(&dep).expect("debug")
                        );
                        return Err(err);
                    }
                }
                Ok(())
            }
            REPLICATED_TASK_NAME => {
                // Scale jobs
                let (jobreq, _) = batch::Job::read_namespaced_job(
                    self.instance_name.as_str(),
                    ns,
                    Default::default(),
                )?;
                let jobres = client.request(jobreq);
                if jobres.is_ok() {
                    let original: batch::Job = jobres.unwrap();
                    let new_job = self.scale_job(original);

                    let (req2, _) = batch::Job::replace_namespaced_job(
                        self.instance_name.as_str(),
                        ns,
                        &new_job,
                        Default::default(),
                    )?;
                    let res2: Result<serde_json::Value, failure::Error> = client.request(req2);
                    if res2.is_err() {
                        let err = res2.unwrap_err();
                        error!(
                            "Scaling error: {}",
                            serde_json::to_string_pretty(&new_job).expect("debug")
                        );
                        return Err(err);
                    }
                }
                Ok(())
            }
            _ => {
                info!("Unsupported workload type: {}", self.workload_type.as_str());
                Ok(())
            }
        }
    }

    /// Scale a deployment
    ///
    /// This takes a base deployment and returns a new deployment with the replica count set.
    pub fn scale_deployment(&self, deployment: apps::Deployment) -> apps::Deployment {
        let new_spec = apps::DeploymentSpec {
            replicas: Some(self.replica_count),
            ..deployment.spec.unwrap()
        };

        apps::Deployment {
            spec: Some(new_spec),
            metadata: deployment.metadata.clone(),
            ..Default::default()
        }
    }

    pub fn scale_job(&self, job: batch::Job) -> batch::Job {
        batch::Job {
            spec: Some(batch::JobSpec {
                parallelism: Some(self.replica_count),
                ..job.spec.unwrap()
            }),
            metadata: job.metadata.clone(),
            ..Default::default()
        }
    }
}

impl TraitImplementation for ManualScaler {
    fn add(&self, ns: &str, client: APIClient) -> TraitResult {
        self.scale(ns, client)
    }
    fn modify(&self, ns: &str, client: APIClient) -> TraitResult {
        self.scale(ns, client)
    }
    fn supports_workload_type(name: &str) -> bool {
        // Only support replicated service and task right now.
        name == REPLICATED_SERVICE_NAME || name == REPLICATED_TASK_NAME
    }
}
