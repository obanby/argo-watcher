// The event log I ould like to produce

"reason": "ResourceCreated",
"message": "Unknown user created application",
involvedObject.kind = "Application"


We will have three different options in this small app

1- Is to get events every 20s for a given app and only print the new changes
2- Is to watch the events changes stream and pretty print some of the info there
3- The store tracking stuff