import { FormControl, FormLabel, Input, FormErrorMessage, FormHelperText } from '@chakra-ui/react';

const InputText = ({
  type,
  value,
  name,
  label,
  onChange,
  disabled,
  meta,
  required,
  helperMessage,
  errorMessage,
}) => {
  const isInvalid = (meta && meta.touched && meta.error) || errorMessage;

  return (
    <FormControl isInvalid={isInvalid} isRequired={required}>
      <FormLabel>{label}</FormLabel>
      <Input
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        bg={'white'}
        disabled={disabled}
        required={required}
      />
      {!isInvalid && helperMessage && <FormHelperText>{helperMessage}</FormHelperText>}
      {isInvalid && <FormErrorMessage>{meta.error || errorMessage}</FormErrorMessage>}
    </FormControl>
  );
};

InputText.defaultProps = {
  isRequired: false,
  value: '',
};

export default InputText;
