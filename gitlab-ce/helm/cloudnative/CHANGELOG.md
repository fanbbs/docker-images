**Note:** This file is automatically generated. Please see the [developer
documentation](doc/development/changelog.md) for instructions on adding your own
entry.

## 2018-06-12

### Removed (1 change)

- Removed version number from statefulset metadata. !318

### Fixed (1 change)

- Resolve No longer possible to disable nginx ingress. !311

### Changed (1 change, 1 of them is from the community)

- Use .Release.Name in the name of all Secrets. !302 (Corey O'Brien)

### Added (3 changes)

- Unicorn: add support for object storage of job artifacts. !255
- Add global setting for imagePullPolicy. !305
- Sidekiq: add object storage configuration. !314

## 2018-06-05

### Changed (1 change)

- GitLab image tags can be set using a global, and default to the chart appVersion. !304

### Added (1 change, 1 of them is from the community)

- Increase minimum repliacs to 2+ via HPAs where relevant. !285 (Corey O'Brien)

### Other (2 changes, 1 of them is from the community)

- Remove placeholder charts. !301 (Corey O'Brien)
- Use upstream gitlab-runner.

## 2018-05-29

### Fixed (1 change, 1 of them is from the community)

- Fix selectors for minio PDB to exclude Job Pods. !292 (Corey O'Brien)

### Added (3 changes, 2 of them are from the community)

- Implement global configuration for Gitaly. !287
- Add options to set annotations on Services. !298 (Corey O'Brien)
- Added soft host anti-affinity by default. (Corey O'Brien)

## 2018-05-22

### Added (3 changes, 3 of them are from the community)

- Add a script for running helm without tiller running in the cluster. !281 (Corey O'Brien)
- Add configuration for auxiliary cron jobs. !283 (Dave Konopka)
- Add basic PodDisruptionBudgets. !284 (Corey O'Brien)

## 2018-05-15

### Changed (2 changes)

- Replace omnibus postgresql instance with postegresql chart based on upstream. !216
- Switch Gitaly Deployment to a StatefulSet. !271

### Added (5 changes, 1 of them is from the community)

- Add prometheus exporter to postgres. !239
- Added configuration options for workhorse. !269 (Corey O'Brien)
- Add ability to configure Redis via Globals. !273
- Add prometheus exporter to gitaly. !274
- Add configuration for ldap authentication.

## 2018-05-08

### Security (1 change)

- Update NGINX Ingress to 0.14.0. !249

### Added (1 change)

- Add configuration options for time_zone, rack_attack, trusted_proxies, and extras. !260

## 2018-05-01

### Added (2 changes, 1 of them is from the community)

- Add configuration options for outgoing email persona. !238 (Corey O'Brien)
- Add object storage support for Uploads. !250

## 2018-04-17

### Fixed (1 change)

- Add updated sidekiq queues for GitLab 10.7. !212

### Added (3 changes, 1 of them is from the community)

- Implement global to configure extnernal static IP to simplify. !206
- Unicorn: add support for configuring Omniauth. !210
- Add Prometheus metrics exporters to Redis and Redis-HA charts. !215 (Dave Konopka)

### Other (2 changes, 1 of them is from the community)

- Bring CHANGELOG logic from gitlab-ce to helm.gitlab.io. !200 (Jason Plum)
- Changelog: introduce changelog_manager, with CI. !208

## 2018-03-22 Alpha

- See [alpha documentation](https://gitlab.com/charts/helm.gitlab.io/blob/master/doc/architecture/alpha.md)
