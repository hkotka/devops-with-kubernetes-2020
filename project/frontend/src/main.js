import App from './App.svelte';
import TodoForm from './TodoForm.svelte';
import TodoList from './TodoList.svelte';

const app = new App({
	target: document.body,
	props: {
		title: 'ToDo Project app'
	}
});

const todo = new TodoForm({
	target: document.body,
	props: {
		txtPlaceholder: 'New ToDo task'
	}
});

const todolist = new TodoList({
	target: document.body
});

//const pongcount = new Pongcount({
//	target: document.body
//})

export default app;