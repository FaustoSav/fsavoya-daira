import axios from 'axios';
export const postOperation = (operation, result) => {
  axios.post('http://localhost:9090/history', {
    operation: operation.toString(),
    result: result.toString()
  })
  .then((response) => {
    console.log(response);
  })
  .catch((error) => {
    console.error('Error al realizar el post:', error);
    throw new Error('Error al guardar la operación en el historial');
  });
};

