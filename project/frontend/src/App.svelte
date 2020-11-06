<script>
	import TodoForm from "./TodoForm.svelte";
	import TodoList from "./TodoList.svelte";
	export let title;
	let apiUrl = "/todos";
</script>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
		overflow: clip;
	}

	h1 {
		color: #000000;
		text-transform: uppercase;
		font-size: 3em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>

<main>
	<h1>{title}</h1>
	<TodoForm txtPlaceholder="New ToDo task" {apiUrl} />
	<TodoList {apiUrl} />
</main>
