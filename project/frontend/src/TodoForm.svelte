<script>
	export let txtPlaceholder;

	let newTodo;
	async function doPost() {
		const res = await fetch("http://localhost:30000/todos", {
			method: "POST",
			body: JSON.stringify({
				name: newTodo,
				done: false,
			}),
		});
		const json = await res.json();
		resetTextInput();
		getTodos;
	}

	function resetTextInput() {
		document.getElementById("todo").value = "";
	}
</script>

<style>
	input[type="text"] {
		width: 50%;
		font-weight: 50;
	}
	button {
		color: black;
		background-color: gold;
		text-transform: uppercase;
		font-weight: 50;
	}
	[id="add-todo"] {
		align-content: center;
	}
	div {
		text-align: center;
	}
</style>

<div id="add-todo">
	<form action="/todos" method="post" id="todo-form">
		<input
			type="text"
			placeholder={txtPlaceholder}
			id="todo"
			name="todo"
			required
			minlength="2"
			maxlength="140"
			bind:value={newTodo} />
		<button type="button" on:click={doPost}>Add ToDo</button>
	</form>
</div>
