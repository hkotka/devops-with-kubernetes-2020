<script>
	const getTodos = (async () => {
		const response = await fetch("https://localhost:30443/todos", {
			method: "GET",
			mode: "no-cors",
		});
		let todos = await response.json();
		console.log(todos);
		return todos;
	})();
</script>

<style>
	div {
		text-align: center;
		align-items: center;
		margin-left: auto;
		margin-right: auto;
	}
	p {
		border: 1px black;
		border-style: groove;
		border-radius: 5px;
		padding: 10px;
		width: 200pt;
		fill: azure;
		font-weight: 50;
		text-align: left;
		align-items: center;
		margin-left: auto;
		margin-right: auto;
	}
	span {
		float: right;
	}
</style>

{#await getTodos}
	<p>Getting tasks...</p>
{:then data}
	{#each data.todos as todo}
		<div>
			<p>
				{todo.name}
				<span>{#if !todo.done}<input type="checkbox" />{/if}</span>
			</p>
		</div>
	{/each}
{:catch error}
	<p>An error ocurred: {error}</p>
{/await}
