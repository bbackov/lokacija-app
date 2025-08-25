import * as Location from 'expo-location';

export async function dohvatiTrenutnuLokaciju() {
  try {

    const { status } = await Location.requestForegroundPermissionsAsync();
    if (status !== 'granted') {
      throw new Error('Dozvola za lokaciju nije odobrena');
    }

 
    const lokacija = await Location.getCurrentPositionAsync({
      accuracy: Location.Accuracy.High,
    });
    console.log("jel tu puca");
    return {
      geografska_sirina: lokacija.coords.latitude,
      geografska_duzina: lokacija.coords.longitude,
      pravac:lokacija.coords.heading,
      preciznost:lokacija.coords.accuracy,
      visina:lokacija.coords.altitude,
      brzina:lokacija.coords.speed,
    };
  } catch (err) {
    throw err;
  }
}