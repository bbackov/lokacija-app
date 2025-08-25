import { StyleSheet, Dimensions } from 'react-native';
const { width, height } = Dimensions.get('window');

export default StyleSheet.create({

    modalOverlay: {
        flex: 1,
        backgroundColor: 'rgba(0,0,0,0.5)',
        justifyContent: 'center',
        alignItems: 'center',
      },
      modalContent: {
        width: '90%',
        backgroundColor: '#fff',
        padding: 20,
        borderRadius: 12,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.3,
        shadowRadius: 6,
        elevation: 5,
      },

      container: {
        padding: width * 0.05,
        marginTop: height * 0.1,
        backgroundColor: '#e5e8ea', 
        flex: 1,
        alignItems: 'center',
      },
    
      tekst: {
        marginBottom: height * 0.015,
        fontSize: 18,
        fontWeight: '500',
        color: '#333',
      },
    
      input: {
        backgroundColor: '#ecf0f1',
        borderWidth: 1,
        borderColor: '#b0bec5',
        paddingHorizontal: width * 0.04,
        paddingVertical: height * 0.015,
        borderRadius: 12,
        fontSize: 16,
        marginBottom: height * 0.025,
        color: '#000',
      },
    
      gumb: {
        borderRadius: 15,
        paddingHorizontal: width * 0.12,
        paddingVertical: height * 0.015,
        marginTop: height * 0.01,
        marginBottom: height * 0.08,
        backgroundColor: '#6c88a6',
        borderColor: '#b0bec5',
        borderWidth: 1,
        alignItems: 'center',
        alignSelf: 'center',
      },
    
      gumbTekst: {
        color: '#fff',
        padding: width * 0.003,
        fontSize: 15,
        fontWeight: '700',
      },

})