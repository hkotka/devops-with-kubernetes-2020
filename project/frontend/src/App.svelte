<script>
	import { onMount } from "svelte";
	import { todoList } from "./store.js";
	import TodoForm from "./TodoForm.svelte";
	import TodoList from "./TodoList.svelte";
	export let title;
	let apiUrl = "http://localhost:8080/todos";

	onMount(async () => {
		let todos;
		const res = await fetch(apiUrl);
		if (res.ok) {
			todos = await res.json();
			todoList.set(todos);
			console.log(todos);
		} else {
			console.log("HTTP-Error: " + res.status);
		}
	});

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
		console.log(newlist);
		todoList.set(newlist);
	}
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
	<TodoForm
		txtPlaceholder="New ToDo task"
		{apiUrl}
		on:notify={updateTodoList} />
	<TodoList />
</main>
