def mosaic(nx, ny, width, height)
  cmd = 'gst-launch-1.0 -e videomixer name=mix'
  ny.times do |y|
    nx.times do |x|
      n = x + y * nx
      cmd << " sink_#{n}::xpos=#{x*width} sink_#{n}::ypos=#{y*height}"
    end
  end
  cmd << " ! videoscale ! video/x-raw,width=#{nx*width},height=#{ny*height}"
  cmd << " ! autovideosink"

  ny.times do |y|
    nx.times do |x|
      n = x + y * nx
      pattern = n % 24
      cmd <<  " videotestsrc pattern=#{pattern} ! video/x-raw, width=#{width}, height=#{height} ! mix.sink_#{n}"
    end
  end

  cmd
end

cmd = mosaic(6, 4, 400, 300)
puts cmd

`#{cmd}`