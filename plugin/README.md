# Plugin Kong

É possível adicionar plugins customizados, feitos em Go, no Kong para realizar ajustes tanto na request quanto no response das APIs.

## O que o plugin faz

Este plugin `kong-tigre-header` adiciona novos headers tanto na request quanto no response das APIs que forem configuradas com ele.

Para compilar o plugin, foi utilizado os seguintes comandos:

```shell
go env -w GOOS=linux
go build -o kong-tigre-header
```

## Como configurar no Kong

Este plugin precisa ter os seguintes valores configurados no `kong.yml`:

```yaml
  - name: usuario-api-1
    url: http://usuario-api-1:80
    routes:
      - name: usuario-api-1-route
        paths:
        - /api1
        strip_path: true

    # CONFIGURAÇÃO DO PLUGIN
    plugins:
      - name: kong-tigre-header ## Nome do plugin
        config:
          request_header: REQUEST ## Parâmetro esperado pelo plugin
          response_header: RESPONSE ## Parâmetro esperado pelo plugin
```

Logo em seguida, é preciso registrar o plugin no `docker-compose` do KONG.

```yaml
volumes:
    - ./plugin/kong-tigre-header:/usr/local/bin/kong-tigre-header
environment:
    - KONG_PLUGINS=bundled,kong-tigre-header
    - KONG_PLUGINSERVER_NAMES=kong-tigre-header
    - KONG_PLUGINSERVER_GO_HEADER_PLUGIN_START_CMD=/usr/local/bin/kong-tigre-header -dump
    - KONG_PLUGINSERVER_GO_HEADER_PLUGIN_QUERY_CMD=/usr/local/bin/kong-tigre-header -dump
```