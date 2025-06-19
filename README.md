# nextcloud-config-generator

> Generate configuration files for Nextcloud.

A CLI tool which generates [configuration files](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html) for Nextcloud
on the fly. Suitable for Kubernetes deployments where the [official Helm Chart](https://github.com/nextcloud/helm/tree/main/charts/nextcloud) seems unfitting or overwhelming.

## Usage

```
Usage:
  nc-cfg-gen [command]

Config Generation
  dynamic     Generates the dynamic configuration, which may change during starts.
  secrets     Generates the secret values instanceid, passwordsalt and secret.

Use "nc-cfg-gen [command] --help" for more information about a command.

```

## Configuration

| Environment Variable | Description | Default | Documentation |
| ------------ | ------------ | ------------ | ------------ |
| `NC_HOST` | The primary public-facing URL |  |  |
| `NC_SCHEME` |  | https |  |
| `NC_TRUSTED_DOMAINS` | All URLs the nextcloud instance will be accessible at |  |  |
| `NC_DB_TYPE` |  |  | [`dbtype`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbtype) |
| `NC_DB_HOST` |  |  | [`dbhost`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbhost) |
| `NC_DB_NAME` |  |  | [`dbname`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbname) |
| `NC_DB_USERNAME` |  |  | [`dbuser`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbuser) |
| `NC_DB_PASSWORD` |  |  | [`dbpassword`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbpassword) |
| `NC_DB_PREFIX` |  | nc_ | [`dbtableprefix`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbtableprefix) |
| `NC_DB_REPLICAS` |  |  | [`dbreplica`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#dbreplica) |
| `NC_MAIL_DOMAIN` |  |  | [`mail_domain`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_domain) |
| `NC_MAIL_FROM_ADDRESS` |  |  | [`mail_from_address`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_from_address) |
| `NC_MAIL_MODE` |  | smtp | [`mail_smtpmode`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtpmode) |
| `NC_MAIL_HOST` |  |  | [`mail_smtphost`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtphost) |
| `NC_MAIL_PORT` |  | 25 | [`mail_smtpport`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtpport) |
| `NC_MAIL_SECURE` |  | true | [`mail_smtpsecure`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtpsecure) |
| `NC_MAIL_USERNAME` |  |  | [`mail_smtpname`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtpname) |
| `NC_MAIL_PASSWORD` |  |  | [`mail_smtppassword`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#mail_smtppassword) |
| `NC_REDIS_ENABLED` |  | false | [`redis`](https://docs.nextcloud.com/server/latest/admin_manual/configuration_server/config_sample_php_parameters.html#redis) |
| `NC_REDIS_HOST` |  |  |  |
| `NC_REDIS_PORT` |  | 6379 |  |
| `NC_REDIS_TIMEOUT` |  |  |  |
| `NC_REDIS_USERNAME` |  |  |  |
| `NC_REDIS_PASSWORD` |  |  |  |
| `NC_REDIS_DATABASE` |  | 0 |  |


## License

[MIT](LICENSE)
