const LOCAL_IP = 'http://192.168.1.104:8085';     
const NGROK_URL = 'https://8d28-88-207-74-244.ngrok-free.app';  


const USE_NGROK = false;

export const BASE_URL = USE_NGROK ? NGROK_URL : LOCAL_IP;