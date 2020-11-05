import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		title: 'ToDo Project app'
	}
});

//const pongcount = new Pongcount({
//	target: document.body
//})

export default app;