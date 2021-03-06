---
description: This section describes the configuration parameters and their types for INX-Spammer.
keywords:
- IOTA Node
- Hornet Node
- Faucet
- Configuration
- JSON
- Customize
- Config
- reference
---


# Core Configuration

INX-Spammer uses a JSON standard format as a config file. If you are unsure about JSON syntax, you can find more information in the [official JSON specs](https://www.json.org).

You can change the path of the config file by using the `-c` or `--config` argument while executing `inx-spammer` executable.

For example:
```bash
inx-spammer -c config_defaults.json
```

You can always get the most up-to-date description of the config parameters by running:

```bash
inx-spammer -h --full
```
## <a id="app"></a> 1. Application

| Name            | Description                                                                                            | Type    | Default value |
| --------------- | ------------------------------------------------------------------------------------------------------ | ------- | ------------- |
| checkForUpdates | Whether to check for updates of the application or not                                                 | boolean | true          |
| stopGracePeriod | The maximum time to wait for background processes to finish during shutdown before terminating the app | string  | "5m"          |

Example:

```json
  {
    "app": {
      "checkForUpdates": true,
      "stopGracePeriod": "5m"
    }
  }
```

## <a id="inx"></a> 2. INX

| Name    | Description                            | Type   | Default value    |
| ------- | -------------------------------------- | ------ | ---------------- |
| address | The INX address to which to connect to | string | "localhost:9029" |

Example:

```json
  {
    "inx": {
      "address": "localhost:9029"
    }
  }
```

## <a id="spammer"></a> 3. Spammer

| Name                  | Description                                                                                                 | Type    | Default value                  |
| --------------------- | ----------------------------------------------------------------------------------------------------------- | ------- | ------------------------------ |
| bindAddress           | The bind address on which the Spammer HTTP server listens                                                   | string  | "localhost:9092"               |
| message               | The message to embed within the spam blocks                                                                 | string  | "We are all made of stardust." |
| tag                   | The tag of the block                                                                                        | string  | "HORNET Spammer"               |
| tagSemiLazy           | The tag of the block if the semi-lazy pool is used (uses "tag" if empty)                                    | string  | "HORNET Spammer Semi-Lazy"     |
| cpuMaxUsage           | Workers remains idle for a while when cpu usage gets over this limit (0 = disable)                          | float   | 0.8                            |
| bpsRateLimit          | The blocks per second rate limit for the spammer (0 = no limit)                                             | float   | 0.0                            |
| workers               | The amount of parallel running spammers                                                                     | int     | 0                              |
| autostart             | Automatically start the spammer on startup                                                                  | boolean | false                          |
| nonLazyTipsThreshold  | The maximum amount of tips in the non-lazy tip-pool before the spammer tries to reduce these (0 = always)   | uint    | 0                              |
| semiLazyTipsThreshold | The maximum amount of tips in the semi-lazy tip-pool before the spammer tries to reduce these (0 = disable) | uint    | 30                             |

Example:

```json
  {
    "spammer": {
      "bindAddress": "localhost:9092",
      "message": "We are all made of stardust.",
      "tag": "HORNET Spammer",
      "tagSemiLazy": "HORNET Spammer Semi-Lazy",
      "cpuMaxUsage": 0.8,
      "bpsRateLimit": 0,
      "workers": 0,
      "autostart": false,
      "nonLazyTipsThreshold": 0,
      "semiLazyTipsThreshold": 30
    }
  }
```

## <a id="profiling"></a> 4. Profiling

| Name        | Description                                       | Type    | Default value    |
| ----------- | ------------------------------------------------- | ------- | ---------------- |
| enabled     | Whether the profiling plugin is enabled           | boolean | false            |
| bindAddress | The bind address on which the profiler listens on | string  | "localhost:6060" |

Example:

```json
  {
    "profiling": {
      "enabled": false,
      "bindAddress": "localhost:6060"
    }
  }
```

## <a id="prometheus"></a> 5. Prometheus

| Name            | Description                                                     | Type    | Default value    |
| --------------- | --------------------------------------------------------------- | ------- | ---------------- |
| enabled         | Whether the prometheus plugin is enabled                        | boolean | false            |
| bindAddress     | The bind address on which the Prometheus HTTP server listens on | string  | "localhost:9312" |
| spammerMetrics  | Whether to include the spammer metrics                          | boolean | true             |
| goMetrics       | Whether to include go metrics                                   | boolean | false            |
| processMetrics  | Whether to include process metrics                              | boolean | false            |
| promhttpMetrics | Whether to include promhttp metrics                             | boolean | false            |

Example:

```json
  {
    "prometheus": {
      "enabled": false,
      "bindAddress": "localhost:9312",
      "spammerMetrics": true,
      "goMetrics": false,
      "processMetrics": false,
      "promhttpMetrics": false
    }
  }
```

