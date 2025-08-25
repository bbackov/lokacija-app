import { View, Text, TextInput, Button ,TouchableOpacity} from 'react-native';
import { useState } from 'react';
import styles from './styles';
import { fetchLogin } from '../../src/api/fetchLogin';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { useRouter } from 'expo-router';


export default function Login() {

 
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [lozinka, setLozinka] = useState('');

  const tokenLogin = async () => {
    try {
      const data = await fetchLogin(email, lozinka);
      await AsyncStorage.setItem('token', data.token);
      alert("prijava uspješna");
      router.replace("/mainpage/home"); 
    } catch (err) {
      alert(err.message || 'Neuspjeh pri povezivanju.');
    
  };

    
  }

  return (

    <View style={styles.container}>
      <Text style={styles.tekst}>E-mail:</Text>
      <TextInput
        placeholder="Unesi E-mail"
        value={email}
        onChangeText={setEmail}
        style={styles.input}
        autoCapitalize="none"
        keyboardType="email-address"
        autoCorrect={false}
      />

      <Text style={styles.tekst}>Lozinka:</Text>
      <TextInput
        placeholder="*********"
        value={lozinka}
        onChangeText={setLozinka}
        secureTextEntry
        style={styles.input}
      />

      <TouchableOpacity style={styles.gumb} onPress={tokenLogin}>
          <Text style={styles.gumbTekst}>Prijavi se</Text>
      </TouchableOpacity>

      <View>
        <Text style={styles.tekst}>Nemas profil registriraj se</Text>
        <TouchableOpacity 
          onPress={() => router.replace("/korisnik/register")}
          style={styles.gumb}> 
          <Text style={styles.gumbTekst}>Registracija</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}