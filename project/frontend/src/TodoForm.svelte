<script>
	import { todoList } from "./store.js";
	export let txtPlaceholder;
	export let apiUrl;
	let newTodo;

	function doPost() {
		fetch(apiUrl, {
			method: "POST",
			body: JSON.stringify({
				name: newTodo,
				done: false,
			}),
		});
		updateTodoList();
		resetTextInput();
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
		} else {
			console.log("HTTP-Error: " + response.status);
		}
		todoList.set(newlist);
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
		<button type="submit" on:click={doPost}>Add ToDo</button>
	</form>
</div>
