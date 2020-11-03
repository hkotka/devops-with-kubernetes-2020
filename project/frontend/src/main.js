import App from './App.svelte';
import TodoForm from './TodoForm.svelte';
//import Pongcount from './Pongcount.svelte';

const app = new App({
	target: document.body,
	props: {
		title: 'ToDo Project app'
	}
});

const todo = new TodoForm({
	target: document.body,
	props: {
		btnName: 'Create',
		txtPlaceholder: 'New ToDo task'
	}
});

//const pongcount = new Pongcount({
//	target: document.body
//})

export default app;