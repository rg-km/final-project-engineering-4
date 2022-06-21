import { useState } from 'react';
import { Button, FormControl, FormLabel, Input, InputGroup, InputRightElement } from '@chakra-ui/react';

const InputPassword = ({ value, name, label, onChange }) => {
  const [show, setShow] = useState(false);

  return (
    <FormControl isRequired overflow={'hidden'}>
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
    </FormControl>
  );
};

export default InputPassword;
