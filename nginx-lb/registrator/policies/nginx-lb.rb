docker_container -> (container) {
  return unless overlay_cidr = container.labels['io.kontena.container.overlay_cidr']
  return unless service = container.env['io.kontena.nginx.service']

  ipv4 = overlay_cidr.split('/')[0]

  if tcp_port = container.env['io.kontena.nginx.tcp']
    return {
      "/kontena/nginx/services/#{service}/tcp" => { port: tcp_port },
      "/kontena/nginx/services/#{service}/backends/#{container.hostname}" => { ipv4: ipv4, tcp: tcp_port },
    }
  elsif udp_port = container.env['io.kontena.nginx.udp']
    return {
      "/kontena/nginx/services/#{service}/udp" => { port: udp_port },
      "/kontena/nginx/services/#{service}/backends/#{container.hostname}" => { ipv4: ipv4, tcp: udp_port },
    }
  else
    return nil
  end
}
