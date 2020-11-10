<script>
	import { todoList } from "./store.js";
	import { fade } from 'svelte/transition';
	$: todos = $todoList;
</script>

<style>
	div {
		text-align: center;
		align-items: center;
		margin-left: auto;
		margin-right: auto;
		width: 40%;
	}
	p {
		border: 1px black;
		border-style: groove;
		border-radius: 5px;
		padding: 10px;
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

{#if todos.todos}
	{#each todos.todos as todo}
		<div transition:fade>
			<p>
				{todo.name}
				<span>{#if !todo.done}<input type="checkbox" />{/if}</span>
			</p>
		</div>
	{/each}
{/if}
