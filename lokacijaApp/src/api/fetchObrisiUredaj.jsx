import { BASE_URL } from "./apiConfig";

export async function fetchObrisiUredaj(token,id_uredaj) {

    try{
        const res = await fetch(`${BASE_URL}/obrisi_uredaj?id=${id_uredaj}`, {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${token}`,
            }
          });
              
        const data=await res.json();
    

    if(!res.ok){
        throw new Error(data.poruka || 'Neuspješno brisanje uredaja');
    }
    return data
}catch(err){
    throw(err)
} 
    
}