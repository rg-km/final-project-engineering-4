import { FormControl, FormLabel, Input, FormErrorMessage, FormHelperText } from '@chakra-ui/react';

const InputText = ({ type, value, name, label, onChange, meta, required, helperMessage, errorMessage }) => {
  return (
    <FormControl isRequired={required}>
      <FormLabel>{label}</FormLabel>
      <Input type={type} name={name} value={value} onChange={onChange} bg={'white'} required={required} />
      {!meta && !errorMessage && helperMessage && <FormHelperText>{helperMessage}</FormHelperText>}
      {((meta && meta.touched && meta.error) || errorMessage) && (
        <FormErrorMessage>{errorMessage}</FormErrorMessage>
      )}
    </FormControl>
  );
};

InputText.defaultProps = {
  isRequired: false,
  value: '',
};

export default InputText;
