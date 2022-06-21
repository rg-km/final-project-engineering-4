import { FormControl, FormLabel, Input } from '@chakra-ui/react';
import React from 'react';

const InputText = ({ type, value, name, label, onChange, required }) => {
  return (
    <FormControl isRequired={required}>
      <FormLabel>{label}</FormLabel>
      <Input type={type} name={name} value={value} onChange={onChange} bg={'white'} required={required} />
    </FormControl>
  );
};

InputText.defaultProps = {
  isRequired: false,
  value: '',
};

export default InputText;
