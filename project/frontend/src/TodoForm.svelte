<script>
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();
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
		resetTextInput();
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

<!-- svelte-ignore a11y-autofocus -->
<div id="add-todo">
	<form
		action="/todos"
		on:submit={() => dispatch('notify')}
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
		<button type="submit" on:click={() => dispatch('notify')}>Add ToDo</button>
	</form>
</div>
