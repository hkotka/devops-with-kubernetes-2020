# Choosing database solution

## Google Cloud SQL
Choosing between Google Cloud SQL versus running and managing Postgres inside k8s with persistent volumes comes down to use case. For production use-case, for most cases it's a good idea to go with Googles managed solution as it provides many critical features that you would have to implement and handle yourself on top of just running the db instance inside k8s cluster like data encryption, backups, applying patches, upgrading database versions and more tricky and complicated things like database replication/high-availability and scalability(adding read-replicas). Also basic things like underlying OS and database configuration is battle tested by Google, along with the whole database lifecycle management and tooling are done in best in class manner. Overall this should be the default option for most prodcution cases, unless there are something critical that the managed option cannot do, for example feature wise.

## Postgres on k8s with persistent volumes
Reasons to go with your own Postgres instance inside k8s cluster with persistent volumes could be for easier tooling integration for setting up cost-effective development environments, like we are doing with gke namespace+deployment per branch during development. This is a good use case for going with k8s Postgres as it's easy to integrate with development workflow and CI/CD tools like Kustomize. Interacting with seperate managed db requires different tools to be integrated into workflows. Production wise, I would consider managing your own Postgres to make sense if you need or want complete control of db configuration, tooling stack and replication/HA setup or need more sophisticated features like sharding that the managed offering does not provide.

## My choice for the project

My choice is to run Postgres on k8s and manage it by myself. Reason for this is that the project is not used for production use case, so many of the benefits of managed DB don't provide any benefit for this project. Managing my own Postgres instances fits better with the tooling I'm using to manage k8s and CI/CD already handles spinning up new Postgres instances from new branches etc. very well and quickly, making the iteration times fast. Using managed db would add more complexity and make it slower to spin up new instances for no tangible benefit.