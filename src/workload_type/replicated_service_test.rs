use kube::{client::APIClient, config::Configuration};

use crate::schematic::component::Component;
use crate::workload_type::*;

use std::collections::BTreeMap;

#[test]
fn test_replicated_service_kube_name() {
    let cli = APIClient::new(mock_kube_config());

    let rs = ReplicatedService {
        name: "de".into(),
        component_name: "hydrate".into(),
        instance_name: "dehydrate".into(),
        namespace: "tests".into(),
        definition: Component {
            ..Default::default()
        },
        params: BTreeMap::new(),
        client: cli,
        owner_ref: None,
    };

    assert_eq!("dehydrate", rs.kube_name().as_str());
}

/// This mock builds a KubeConfig that will not be able to make any requests.
fn mock_kube_config() -> Configuration {
    Configuration {
        base_path: ".".into(),
        client: reqwest::Client::new(),
    }
}
