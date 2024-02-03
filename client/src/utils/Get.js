import axios from 'axios';

export const getOperationsHistoy = async () => {
  try {
    const response = await axios.get('http://localhost:9090/history');
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};
