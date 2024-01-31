<script>
	import '../routes/styles.css';
	import Button from './Button.svelte';

	const buttons = [
		{ id: 'clear', customClass: 'btn-dark', value: 'C' },
		{ id: 'open-parenthesis', customClass: 'btn-orange', value: '(' },
		{ id: 'close-parenthesis', customClass: 'btn-orange', value: ')' },
		{ id: 'divide', customClass: 'btn-orange', value: '/' },
		{ id: 'seven', customClass: 'btn-grey', value: '7' },
		{ id: 'eight', customClass: 'btn-grey', value: '8' },
		{ id: 'nine', customClass: 'btn-grey', value: '9' },
		{ id: 'multiply', customClass: 'btn-orange', value: 'X' },
		{ id: 'four', customClass: 'btn-grey', value: '4' },
		{ id: 'five', customClass: 'btn-grey', value: '5' },
		{ id: 'six', customClass: 'btn-grey', value: '6' },
		{ id: 'subtract', customClass: 'btn-orange', value: '-' },
		{ id: 'one', customClass: 'btn-grey', value: '1' },
		{ id: 'two', customClass: 'btn-grey', value: '2' },
		{ id: 'three', customClass: 'btn-grey', value: '3' },
		{ id: 'add', customClass: 'btn-orange', value: '+' },
		{ id: 'zero', customClass: 'btn-grey col-span-2 w-full', value: '0' },
		{ id: 'decimal', customClass: 'btn-grey', value: '.' },
		{ id: 'equals', customClass: 'btn-green', value: '=' }
	];
</script>

<section class=" calculator-container">
	{#each buttons as button}
		<Button btnInfo={button} />
	{/each}
</section>

<style>
</style>
