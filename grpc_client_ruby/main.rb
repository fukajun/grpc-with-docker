$LOAD_PATH << '../helloworld'
require '../helloworld/helloworld_services_pb'
require '../helloworld/helloworld_pb'

GRPC::Core::ChannelCredentials.constants
stub = Helloworld::Greeter::Stub.new('localhost:5000', :this_channel_is_insecure)
stub.watch(Helloworld::Channel.new({name: Time.now.to_s})).each do |i|
  puts i.name
end
