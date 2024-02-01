<script>
	//IMPORTS
	import '../routes/styles.css';
	import Button from './Button.svelte';
	import { evaluate, prodDependencies, string } from 'mathjs';
	import { Toaster, toast } from 'svelte-sonner';

	//ARRAY DE OBJETOS PARA MAPEAR BOTONES
	const buttons = [
		{ id: 'clear', customClass: 'btn-dark', value: 'C' },
		{ id: 'open-parenthesis', customClass: 'btn-orange', value: '(' },
		{ id: 'close-parenthesis', customClass: 'btn-orange', value: ')' },
		{ id: 'divide', customClass: 'btn-orange', value: '/' },
		{ id: 'seven', customClass: 'btn-grey', value: '7' },
		{ id: 'eight', customClass: 'btn-grey', value: '8' },
		{ id: 'nine', customClass: 'btn-grey', value: '9' },
		{ id: 'multiply', customClass: 'btn-orange', value: '*' },
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

	//ESTADOS
	//Valor iniciar del input
	let inputValue = '';
	//Operacion previa, para no hacer otra vez el calculo si se sigue apretando el = y el input es el mismo
	let previusOperation = '';
	//Resultado
	let result = '';

	//Limpia la operacion
	const clearInput = () => {
		inputValue = '';
		previusOperation = '';
	};
	//Cálculo del resultado
	const calculateResult = () => {
		if (inputValue) {
			try {
				if (inputValue === previusOperation) {
					toast.info('No esperes un resultado diferente si la operación es la misma');
				} else {
					result = evaluate(inputValue);
					previusOperation = inputValue;
				}
			} catch (error) {
				toast.warning('Parece que hay un error de sintaxis');
			}
		} else {
			toast.error('No hay valores para realizar la operación');
		}
	};

	const handleOperation = (value) => {
		if (value == '=') {
			calculateResult();
		} else if (value == 'C') {
			clearInput();
		} else {
			inputValue += value;
		}
	};
</script>

<div class="calculator-container">
	<Toaster richColors position="bottom-center" />
	<section class="operations-container">
		<input
			class="bg-transparent outline-none cursor-default"
			type="text"
			name="input"
			id="cal"
			value={inputValue}
		/>
		<span class="text-3xl text-white tracking-wide">{result}</span>
	</section>
	<section class="buttons-container">
		{#each buttons as button}
			<button
				class="{button.customClass} btn"
				on:click={() => handleOperation(button.value.toString())}>{button.value}</button
			>
		{/each}
	</section>
</div>

<style>
</style>
