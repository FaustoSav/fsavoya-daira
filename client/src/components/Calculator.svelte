<script>
	import '../routes/styles.css';
	import { evaluate } from 'mathjs';
	import { Toaster, toast } from 'svelte-sonner';
	import { postOperation } from '../utils/Post';
	import { Buttons } from '../utils/Buttons';

	let inputValue = '';
	let previusOperation = '';
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
					postOperation(inputValue, result);
				} catch (error) {
					toast.warning('Parece que hay un error de sintaxis');
					console.log(error);
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
		{#each Buttons as button}
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
