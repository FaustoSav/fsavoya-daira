<script>
	//IMPORTS
	import '../routes/styles.css';
	import { evaluate, prodDependencies, string } from 'mathjs';
	import { Toaster, toast } from 'svelte-sonner';

	//Array de objetos botones para mapear
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
	let result = '0';

	//FUNCIONES
	//Limpia la operacion
	const clearInput = () => {
		inputValue = previusOperation = result = '';
	};
	//Cálculo del resultado
	const calculateResult = () => {
		if (inputValue) {
			if (inputValue === previusOperation) {
				toast.info('Recuerda que la operación es la misma. El resultado se mantendrá igual.');
			} else {
				try {
					result = evaluate(inputValue);
					previusOperation = inputValue;
				} catch (error) {
					toast.warning('Parece que hay un error de sintaxis');
				}
			}
		} else {
			toast.error('No hay valores para realizar la operación');
		}
	};

	//Funcion a ejecutar siempre que se clickee un boton
	const handleOperation = (value) => {
		//Si se apreta el signo igual, se hace el calculo y un par de validaciones
		if (value == '=') {
			calculateResult();
		} else if (value == 'C') {
			//Si se apreta para C, se hace un clear de todos los valores
			clearInput();
		} else {
			//Si se apreta cualquier otro boton, simplemente se lo agrega a la cadena para hacer el evaluate de mathjs
			inputValue += value;
		}
	};
</script>

<div class="calculator-container">
	<!-- Mensaje de alerta -->
	<Toaster richColors position="bottom-center" />
	<section class="operations-container">
		<!-- Input donde se va escribiendo la operacion -->
		<input class="input" type="text" name="input" id="cal" value={inputValue} />
		<!-- Resultado de la operacion -->
		<span class="result">{result.toLocaleString()}</span>
	</section>
	<!-- Teclado / Mapeo de botones -->
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
	.calculator-container {
		@apply w-[310px] rounded-3xl border-[1px] border-white border-opacity-10 overflow-hidden;
	}


	.operations-container {
		@apply w-full min-h-28 p-7 flex flex-col justify-center items-end text-xl text-gray-300 gap-2 text-clip bg-black bg-opacity-30;
	}
</style>
