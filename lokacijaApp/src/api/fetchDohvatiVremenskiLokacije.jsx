import { BASE_URL } from "./apiConfig";

export async function fetchDohvatiVremenskiLokaciju(token, id_uredaj, pocetak, kraj) {
    try {
        const res = await fetch(`${BASE_URL}/dohvati_vremenskilokacije?id=${id_uredaj}&pocetak=${pocetak}&kraj=${kraj}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
        });

        const data = await res.json();

        if (!res.ok) {
            throw new Error(data.poruka || 'Neuspješno dohvacanje lokacije.');
        }

        return Array.isArray(data.lokacije) ? data.lokacije : [];

    } catch (err) {
        alert(err.message);
        throw err;
    }
}