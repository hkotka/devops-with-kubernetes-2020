<script>
	export let apiUrl;
	let todoList = [];

	const getTodos = (async () => {
		const response = await fetch(apiUrl, {
			method: "GET",
		});
		if (response.ok) {
			todoList = await response.json();
		} else {
			console.log("HTTP-Error: " + response.status);
		}
		console.log(todoList);
		return todoList;
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
	Getting tasks...
{:then todos}
	{#if todos.todos == null}
		No Todo's
	{:else}
		{#each todos.todos as todo}
			<div>
				<p>
					{todo.name}
					<span>{#if !todo.done}<input type="checkbox" />{/if}</span>
				</p>
			</div>
		{/each}
	{/if}
{:catch error}
	<p>An error ocurred: {error}</p>
{/await}
