import { useEffect, useState } from 'react';
import { Stack, useRouter } from 'expo-router';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { fetchGetKorisnik } from '../src/api/fetchGetKorisnik';

export default function Index() {
  const [provjereno, setProvjereno] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const provjeriToken = async () => {
      const token = await AsyncStorage.getItem("token");

      if (!token) {
        router.replace("/korisnik/login");
        setProvjereno(true);
        return;
      }

      try {
        const korisnik = await fetchGetKorisnik(token);
        if (korisnik) {
          router.replace("/mainpage/home");
        } else {
          
          await AsyncStorage.removeItem("token");
          router.replace("/korisnik/login");
        }
      } catch (err) {
        
        await AsyncStorage.removeItem("token");
        router.replace("/korisnik/login");
      }

      setProvjereno(true);
    };

    provjeriToken();
  }, []);

  if (!provjereno) return null;

  return <Stack.Screen options={{ headerShown: false }} />;
}