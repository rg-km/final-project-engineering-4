import { Button } from '@chakra-ui/button';
import { Checkbox } from '@chakra-ui/checkbox';
import { Heading, Stack } from '@chakra-ui/layout';
import { Table, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/table';
import { useParams } from 'react-router-dom';

function KelasDetail() {
  let { id } = useParams();
  console.log(id);
  
  return (
    <Stack py={8} spacing={4}>
      <Heading fontSize={'3xl'}>Bahasa Indonesia • 9F</Heading>
      <Stack bg={'white'} rounded={'md'} direction={'row'} spacing={8} padding={4}>
        <Stack flex={1}>
          <Button>Absen</Button>
          <Button variant={'ghost'}>Jadwal</Button>
          <Button variant={'ghost'}>Tugas</Button>
        </Stack>
        <Stack width={'5xl'}>
          <Table>
            <Thead>
              <Tr>
                <Th>No</Th>
                <Th>Nama Lengkap</Th>
                <Th>Hari Ke-1</Th>
                <Th>Hari Ke-2</Th>
                <Th>Hari Ke-3</Th>
              </Tr>
            </Thead>
            <Tbody>
              <Tr>
                <Td>1</Td>
                <Td>Budi</Td>
                <Td>
                  <Checkbox disabled defaultChecked />
                </Td>
                <Td>
                  <Checkbox disabled />
                </Td>
                <Td>
                  <Checkbox disabled defaultChecked />
                </Td>
              </Tr>
              <Tr>
                <Td>2</Td>
                <Td>Joko</Td>
                <Td>
                  <Checkbox disabled defaultChecked />
                </Td>
                <Td>
                  <Checkbox disabled />
                </Td>
                <Td>
                  <Checkbox disabled defaultChecked />
                </Td>
              </Tr>
            </Tbody>
          </Table>
        </Stack>
      </Stack>
    </Stack>
  );
}

export default KelasDetail;
