import { View, Text, TextInput, Button ,TouchableOpacity} from 'react-native';
import { Platform } from 'react-native';
import RNPickerSelect from 'react-native-picker-select';
import MapView from 'react-native-maps';
import { useEffect, useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { fetchDodajLokaciju } from '../src/api/fetchDodajLokaciju';
import { dohvatiTrenutnuLokaciju } from './trenutnaLokacija';
import stylesMainPage from '../app/mainpage/stylesMainPage';
import { fetchDohvatiZadnjuLokaciju } from '../src/api/fetchDohvatiZadnjuLokaciju';
import { fetchDohvati10Lokacija } from '../src/api/fetchDohvati10Lokacija';
import { fetchDohvatiVremenskiLokaciju } from '../src/api/fetchDohvatiVremenskiLokacije';
import { Marker } from 'react-native-maps';
import DateTimePickerModal from "react-native-modal-datetime-picker";




export default function Karta(){


    const [izbornik,setIzbornik]=useState('zadnja');
    const [lokacije,setlokacije]=useState([]);
    const [pocetak, setPocetak] = useState(new Date('2025-05-21T00:00:00Z'));
    const [kraj, setKraj] = useState(new Date());
    const [showPickerOd, setShowPickerOd] = useState(false);
    const [showPickerDo, setShowPickerDo] = useState(false);
    const [fullscreen,setFullscreen]=useState(false);


      const handleConfirmOd = (date) => {
        setPocetak(date);
        setShowPickerOd(false)
      };

      const handleConfirmDo = (date) => {
        setKraj(date);
        setShowPickerDo(false)
      };

    useEffect(() => {
        const intervalDohvat = setInterval(async () => {
          try {
            const token = await AsyncStorage.getItem('token');
            const id = await AsyncStorage.getItem('id_uredaj');
      
            if (!token || !id) return;
      
            if (izbornik === "zadnja") {
              const lokacija = await fetchDohvatiZadnjuLokaciju(token, id);
              console.log(lokacija)
              if (
                lokacija &&
                typeof lokacija.geografska_sirina === 'number' &&
                typeof lokacija.geografska_duzina === 'number'
              ) {
                setlokacije([lokacija]);
              }else {
                setlokacije([]); 
              }
              
            } else if (izbornik === "zadnjih10") {
              const data = await fetchDohvati10Lokacija(token, id);
              const validne = data.filter(
                l =>
                  l &&
                  typeof l.geografska_sirina === 'number' &&
                  typeof l.geografska_duzina === 'number'
              );
              setlokacije(validne);
            }
      
          } catch (err) {
            console.log('Greška dohvaćanje lokacije:', err);
          }
        }, 10000);
      
        return () => clearInterval(intervalDohvat);
      }, [izbornik]);

      useEffect(() => {
        const vremenskiLokacije = async () => {
            if (izbornik !== 'vremenski') return;

            try {
                const token = await AsyncStorage.getItem('token');
                const id = await AsyncStorage.getItem('id_uredaj');
                const data = await fetchDohvatiVremenskiLokaciju(token, id, pocetak.toISOString(), kraj.toISOString());
                setlokacije(Array.isArray(data)? data :[]);
            } catch (err) {
                console.log('Greška  dohvaćanje lokacije vremenski ', err);
            }
        }
        vremenskiLokacije();
    }, [izbornik, pocetak, kraj]);


    


    useEffect(()=>{
        const intervalLokacija=setInterval(async() => {
            
            try{
                const token=await AsyncStorage.getItem('token');
                const lokacija=await dohvatiTrenutnuLokaciju();
                const id = await AsyncStorage.getItem('id_uredaj');
                await fetchDodajLokaciju(
                    token,
                    lokacija.geografska_sirina,
                    lokacija.geografska_duzina,
                    lokacija.pravac,
                    lokacija.preciznost,
                    lokacija.visina,
                    lokacija.brzina,
                    parseInt(id)
                  );
            }catch(err){
                console.log('Greška prilikom slanje lokacije:', err);
            }
        }, 10000);

        return ()=>clearInterval(intervalLokacija);
    },[]);

    useEffect(() => {
        setlokacije([]); 
      }, [izbornik]);

    return(
    <View style={stylesMainPage.lokacijaBlok}>
        <View style={stylesMainPage.izbornik}>
            <Text style={stylesMainPage.tekst}>Izbornik:</Text>
            <RNPickerSelect
                useNativeAndroidPickerStyle={false}
                value={izbornik}
                style={{
                    inputIOS: stylesMainPage.select,
                    inputAndroid: stylesMainPage.select,
                    viewContainer: stylesMainPage.select,
                    placeholder: { color: '#999' },
                }} 
                onValueChange={(value)=>setIzbornik(value)}
                items={[
                    {label:"Zadnja lokacija",value:"zadnja"},
                    {label:"Zadnjih 10 lokacija",value:"zadnjih10"},
                    {label:"Vremenski lokacija",value:"vremenski"}
                ]}
            />
        </View>

        {izbornik==='vremenski'&&
        <View style={{ flexDirection: 'row', justifyContent: 'space-around', marginBottom: 10 }}>
            <TouchableOpacity onPress={() => setShowPickerOd(true)} style={stylesMainPage.dateButton}>
                <Text style={stylesMainPage.dateButtonText}> Od: {pocetak.toLocaleString()}</Text>
            </TouchableOpacity>
            
            <TouchableOpacity onPress={() => setShowPickerDo(true)} style={stylesMainPage.dateButton}>
                <Text style={stylesMainPage.dateButtonText}> Od: {kraj.toLocaleString()}</Text>
            </TouchableOpacity>
    
        </View>}

       

        <DateTimePickerModal
            isVisible={showPickerOd}
            mode="datetime"
            is24Hour={true}
            onConfirm={handleConfirmOd}
            onCancel={()=>setShowPickerOd(false)}
        />

        <DateTimePickerModal
            isVisible={showPickerDo}
            mode="datetime"
            is24Hour={true}
            onConfirm={handleConfirmDo}
            onCancel={()=>setShowPickerDo(false)}
        />

        <TouchableOpacity  onPress={() => setFullscreen(true)} activeOpacity={0.9}>
        <View style={stylesMainPage.karta}>
         <MapView
                style={{ width: '100%', height: 300, borderRadius: 12 }}
                initialRegion={{
                    latitude: 45.815399,
                    longitude: 15.966568,
                    latitudeDelta: 0.05,
                    longitudeDelta: 0.05,
                }}
        >
            {lokacije.map((lokacija, index) => (
                <Marker
                    key={`${lokacija.id_lokacije}-${index}`}
                    coordinate={{
                    latitude: lokacija.geografska_sirina,
                    longitude: lokacija.geografska_duzina
                    }}
                    title={`Lokacija ${index + 1}`}
                    description={new Date(lokacija.vrijeme).toLocaleString()}
                />
                ))}
            </MapView>
        </View>
        </TouchableOpacity>

        {fullscreen&&(
            <View style={stylesMainPage.fullscreenMapContainer}>
                <TouchableOpacity onPress={()=>setFullscreen(false)}
                    style={stylesMainPage.zatvoriGumb}>
                    <Text style={stylesMainPage.zatvoriTekst}>Zatvori</Text>
                </TouchableOpacity>
                <MapView
                style={{ width: '100%', height: 300, borderRadius: 12 }}
                initialRegion={{
                    latitude: 45.815399,
                    longitude: 15.966568,
                    latitudeDelta: 0.05,
                    longitudeDelta: 0.05,
                }}
        >
                {lokacije.map((lokacija, index) => (
                <Marker
                    key={`${lokacija.id_lokacije}-${index}`}
                    coordinate={{
                    latitude: lokacija.geografska_sirina,
                    longitude: lokacija.geografska_duzina
                    }}
                    title={`Lokacija ${index + 1}`}
                    description={new Date(lokacija.vrijeme).toLocaleString()}
                />
                ))}
            </MapView>
            </View>
        )}

    </View>
    );
}