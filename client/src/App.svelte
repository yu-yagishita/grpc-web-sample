<script>
	import { GreeterPromiseClient } from "./helloworld/proto/helloworld_grpc_web_pb"
	import { HelloRequest } from "./helloworld/proto/helloworld_pb";

	export let name;
	// gRPCの呼び出しのテストをする。
	let resp = '';
	let grpcTestName = 'wryyy';
	const client = new GreeterPromiseClient('http://localhost:8080');
	const handleClick = () => {
		const request = new HelloRequest();
		request.setName(grpcTestName);
		console.log(grpcTestName);
		client.sayHello(request).then((reply) => {
		console.log(reply.getMessage())
		resp = reply.getMessage();
		});
		grpcTestName += 'y';
	};
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<button on:click="{handleClick}">gRPCのテストボタン</button>
	<p>{resp}</p>

</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
