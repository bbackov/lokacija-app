import { BASE_URL } from "./apiConfig";

export async function fetchDohvati10Lokacija(token,id_uredaj) {

    try{
        const res=await fetch(`${BASE_URL}/dohvati_10lokacija?id=${id_uredaj}`,{
            method:'GET',
            headers:{   'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
        });

        const data=await res.json();
        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješno dohvacanje lokacije.');
        }
        return data.lokacije;
        }catch(err){
            alert(err.message);
            throw err
        }
    
}