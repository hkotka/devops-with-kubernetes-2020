<script>
	export let txtPlaceholder;
	let newTodo;

	async function doPost() {
		const res = await fetch("/todos", {
			method: "POST",
			body: JSON.stringify({
				name: newTodo,
				done: false,
			}),
		});
		const json = await res.json();
		resetTextInput();
		window.location.href = window.location.href;
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
	<form
		action="/todos"
		method="post"
		on:submit|preventDefault={doPost}
		id="todo-form">
		<input
			type="text"
			placeholder={txtPlaceholder}
			id="todo"
			name="todo"
			required
			minlength="2"
			maxlength="140"
			bind:value={newTodo} />
		<button type="submit">Add ToDo</button>
	</form>
</div>
