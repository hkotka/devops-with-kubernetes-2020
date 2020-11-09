<script>
	import { todoList } from "./store.js";
	export let txtPlaceholder;
	export let apiUrl;
	let newTodo;

	async function doPost() {
		const response = await fetch(apiUrl, {
			method: "POST",
			body: JSON.stringify({
				name: newTodo,
				done: false,
			}),
		});
		if (response.ok) {
			resetTextInput();
			updateTodoList();
		} else {
			console.log("HTTP-Error: " + response.status);
		}
	}

	function resetTextInput() {
		document.getElementById("todo").value = "";
	}

	async function updateTodoList() {
		let newlist;
		const response = await fetch(apiUrl, {
			method: "GET",
		});
		if (response.ok) {
			newlist = await response.json();
			todoList.set(newlist);
			console.log(newlist);
		} else {
			console.log("HTTP-Error: " + response.status);
		}
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

<!-- svelte-ignore a11y-autofocus -->
<div id="add-todo">
	<form
		action="/todos"
		method="post"
		on:submit|preventDefault={doPost}
		id="todo-form">
		<input
			type="text"
			placeholder={txtPlaceholder}
			autofocus="true"
			id="todo"
			name="todo-input"
			required
			minlength="2"
			maxlength="140"
			bind:value={newTodo} />
		<button type="submit">Add ToDo</button>
	</form>
</div>
