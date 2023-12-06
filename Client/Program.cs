using System.Threading.Tasks;
using Grpc.Net.Client;
using GrpcGreeterClient;

namespace Client
{
    internal class Program
    {
        static async Task Main(string[] args)
        {
            using var channel = GrpcChannel.ForAddress("http://localhost:50051");
            var client = new Greeter.GreeterClient(channel);
            var reply = await client.SayHelloAsync(
                  new HelloRequest { Name = "GreeterClient" });
            Console.WriteLine("Greeting: " + reply.Message);
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}
