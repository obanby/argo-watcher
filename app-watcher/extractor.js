let obj = {
    "metadata": {
        "resourceVersion": "18375"
    },
    "items": [
        {
            "metadata": {
                "name": "sock-shop-example.16a1a8996c3c919c",
                "namespace": "argocd",
                "uid": "f75fadf7-d540-4535-b4ac-17591451c5a7",
                "resourceVersion": "15202",
                "creationTimestamp": "2021-09-04T15:38:36Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:38:36Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15201"
            },
            "reason": "ResourceCreated",
            "message": "Unknown user created application",
            "source": {
                "component": "argocd-server"
            },
            "firstTimestamp": "2021-09-04T15:38:36Z",
            "lastTimestamp": "2021-09-04T15:38:36Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a899c8ef7c35",
                "namespace": "argocd",
                "uid": "1873af3c-0ac0-4bd1-89f4-403151c675ed",
                "resourceVersion": "15205",
                "creationTimestamp": "2021-09-04T15:38:38Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:38:38Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15201"
            },
            "reason": "ResourceUpdated",
            "message": "Updated sync status:  -> OutOfSync",
            "source": {
                "component": "argocd-application-controller"
            },
            "firstTimestamp": "2021-09-04T15:38:38Z",
            "lastTimestamp": "2021-09-04T15:38:38Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a899c946937a",
                "namespace": "argocd",
                "uid": "4aa331a3-2d5a-4055-b80f-b4aa36ea0b44",
                "resourceVersion": "15206",
                "creationTimestamp": "2021-09-04T15:38:38Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:38:38Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15201"
            },
            "reason": "ResourceUpdated",
            "message": "Updated health status:  -> Missing",
            "source": {
                "component": "argocd-application-controller"
            },
            "firstTimestamp": "2021-09-04T15:38:38Z",
            "lastTimestamp": "2021-09-04T15:38:38Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a89f512f4e14",
                "namespace": "argocd",
                "uid": "25d6d020-ed79-441a-9b78-8134bf4d57c9",
                "resourceVersion": "15256",
                "creationTimestamp": "2021-09-04T15:39:02Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:39:02Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15255"
            },
            "reason": "OperationStarted",
            "message": "Unknown user initiated sync to HEAD (5a4ab415d85419516870bada4afeb87b109a0c00)",
            "source": {
                "component": "argocd-server"
            },
            "firstTimestamp": "2021-09-04T15:39:02Z",
            "lastTimestamp": "2021-09-04T15:39:02Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a89fd872eecf",
                "namespace": "argocd",
                "uid": "217213a1-23b4-4c9f-b813-95a6defdedc9",
                "resourceVersion": "15376",
                "creationTimestamp": "2021-09-04T15:39:04Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:39:04Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15255"
            },
            "reason": "OperationCompleted",
            "message": "Sync operation to 5a4ab415d85419516870bada4afeb87b109a0c00 succeeded",
            "source": {
                "component": "argocd-application-controller"
            },
            "firstTimestamp": "2021-09-04T15:39:04Z",
            "lastTimestamp": "2021-09-04T15:39:04Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a89fe03bb3b4",
                "namespace": "argocd",
                "uid": "119d9aef-a352-48d6-9b01-83806b1d8fd7",
                "resourceVersion": "15406",
                "creationTimestamp": "2021-09-04T15:39:04Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:39:04Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15334"
            },
            "reason": "ResourceUpdated",
            "message": "Updated sync status: OutOfSync -> Synced",
            "source": {
                "component": "argocd-application-controller"
            },
            "firstTimestamp": "2021-09-04T15:39:04Z",
            "lastTimestamp": "2021-09-04T15:39:04Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        },
        {
            "metadata": {
                "name": "sock-shop-example.16a1a89fe5a73159",
                "namespace": "argocd",
                "uid": "e5f66a8f-e1db-4cdb-b3c9-ed1edc8e2f56",
                "resourceVersion": "15424",
                "creationTimestamp": "2021-09-04T15:39:04Z",
                "annotations": {
                    "dest-namespace": "argoexamples",
                    "dest-server": "https://kubernetes.default.svc"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-09-04T15:39:04Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:count": {},
                            "f:firstTimestamp": {},
                            "f:involvedObject": {
                                "f:apiVersion": {},
                                "f:kind": {},
                                "f:name": {},
                                "f:namespace": {},
                                "f:resourceVersion": {},
                                "f:uid": {}
                            },
                            "f:lastTimestamp": {},
                            "f:message": {},
                            "f:metadata": {
                                "f:annotations": {
                                    ".": {},
                                    "f:dest-namespace": {},
                                    "f:dest-server": {}
                                }
                            },
                            "f:reason": {},
                            "f:source": {
                                "f:component": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "involvedObject": {
                "kind": "Application",
                "namespace": "argocd",
                "name": "sock-shop-example",
                "uid": "5920915b-52cd-4f0f-99e7-f173e4316f0c",
                "apiVersion": "argoproj.io/v1alpha1",
                "resourceVersion": "15334"
            },
            "reason": "ResourceUpdated",
            "message": "Updated health status: Missing -> Progressing",
            "source": {
                "component": "argocd-application-controller"
            },
            "firstTimestamp": "2021-09-04T15:39:04Z",
            "lastTimestamp": "2021-09-04T15:39:04Z",
            "count": 1,
            "type": "Normal",
            "eventTime": null,
            "reportingComponent": "",
            "reportingInstance": ""
        }
    ]
}

obj.items.forEach((item) => console.log(`${item.reason}: ${item.message}`))
