import { BASE_URL } from "./apiConfig";


export async function fetchDodajLokaciju(token,geografska_sirina,geografska_duzina,pravac,preciznost,visina,brzina,id_uredaj){
    console.log("bar sam tu");
    const body = {
        geografska_sirina,
        geografska_duzina,
        pravac,
        preciznost,
        visina,
        brzina,
        id_uredaj
      };
      
      console.log("Šaljem JSON:", JSON.stringify(body));
    try{
        const res=await fetch(`${BASE_URL}/dodaj_lokaciju`,{
            method:'POST',
            headers:{   'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body:JSON.stringify(body)
        });

        const data=await res.json();

        if(!res.ok){
            throw new Error(data.poruka || 'Neuspješno dodavanje lokacije.');
        }
        return data;
        }catch(err){
            alert(err.message);
            throw err
        }
    }
    
