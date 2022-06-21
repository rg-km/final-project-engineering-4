import { useState } from 'react';
import {
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Input,
  InputGroup,
  InputRightElement,
} from '@chakra-ui/react';

const InputPassword = ({ value, name, label, onChange, meta, errorMessage, helperMessage }) => {
  const [show, setShow] = useState(false);

  const isInvalid = (meta && meta.touched && meta.error) || errorMessage;

  return (
    <FormControl isInvalid={isInvalid} isRequired overflow={'hidden'}>
      <FormLabel>{label}</FormLabel>
      <InputGroup>
        <Input
          name={name}
          value={value}
          type={show ? 'text' : 'password'}
          onChange={onChange}
          required
          bg={'white'}
        />
        <InputRightElement width="4.5rem">
          <Button
            h="1.75rem"
            size="sm"
            onClick={() => {
              setShow(!show);
            }}>
            {show ? 'Hide' : 'Show'}
          </Button>
        </InputRightElement>
      </InputGroup>

      {!isInvalid && helperMessage && <FormHelperText>{helperMessage}</FormHelperText>}
      {isInvalid && <FormErrorMessage>{meta.error || errorMessage}</FormErrorMessage>}
    </FormControl>
  );
};

export default InputPassword;
