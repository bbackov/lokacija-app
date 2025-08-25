
import { BASE_URL } from "./apiConfig";

export async function fetchObrisiProfil(token){

    try{
        const res=await fetch(`${BASE_URL}/obrisi_profil`,{
            method:'DELETE',
            headers: {
                'Authorization': `Bearer ${token}`,
              }
        });

        const data=await res.json();
        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješna Odjava.');
        }
        return data;
    }catch(err){
        throw(err);
    }

}