def docker_container(container)
  return unless overlay_cidr = container.labels['io.kontena.container.overlay_cidr']
  return unless service = container.env['io.kontena.nginx.service']

  ipv4 = overlay_cidr.split('/')[0]

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
