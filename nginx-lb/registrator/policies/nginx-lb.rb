def docker_container(container)
  return unless overlay_cidr = container.labels['io.kontena.container.overlay_cidr']
  return unless service = container.env['io.kontena.nginx.service']

  ipv4 = overlay_cidr.split('/')[0]

  if port = container.env['io.kontena.nginx.http']
    aliases = container.env['io.kontena.nginx.http.aliases']
    aliases = aliases.split if aliases

    config = {
      'port'     => container.env['io.kontena.nginx.http.port'],
      'host'     => container.env['io.kontena.nginx.http.host'],
      'aliases'  => aliases,
      'location' => container.env['io.kontena.nginx.http.location'],
    }.reject{|key, value| value.nil?}

    yield ({
      "/kontena/nginx/http/#{service}/server" => config,
      "/kontena/nginx/http/#{service}/upstreams/#{container.hostname}" => { ipv4: ipv4, port: port },
    })
  end

  if port = container.env['io.kontena.nginx.tcp']
    yield ({
      "/kontena/nginx/tcp/#{service}/server" => { port: port },
      "/kontena/nginx/tcp/#{service}/upstreams/#{container.hostname}" => { ipv4: ipv4, port: port },
    })
  end

  if port = container.env['io.kontena.nginx.udp']
    yield ({
      "/kontena/nginx/udp/#{service}/server" => { port: port },
      "/kontena/nginx/udp/#{service}/upstreams/#{container.hostname}" => { ipv4: ipv4, port: port },
    })
  end
end
