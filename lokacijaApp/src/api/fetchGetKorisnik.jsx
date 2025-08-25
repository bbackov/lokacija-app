
import { BASE_URL } from "./apiConfig";

export async function fetchGetKorisnik(token){

    try{
        const res=await fetch(`${BASE_URL}/dohvati_korisnika`,{
            method:'GET',
            headers:{
                'Authorization': `Bearer ${token}`,
            }
        });
    

    const data= await res.json();
    
    if(!res.ok){
        throw new Error(data.poruka || 'Neuspješno dohvaćanje podataka.');
    }
    console.log("Odgovor s backenda:", data);
    return data.korisnik
    }catch(err){
        throw(err);
    }
}