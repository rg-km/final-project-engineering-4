import { Box, Button, Flex, Input, Stack, VStack } from "@chakra-ui/react";
import { PATH } from "@config/path";
import Card from "./components/Card";
import MenuItem from "./components/MenuItem";
import Navbar from "./components/Navbar";

const Class = () => {
    // let data = [
    //     {
    //         id_siswa: 1,

    //     }
    // ]
    return <Box bg={'gray.100'} minH={'100vh'}>
        <Navbar />
        <Box px={{ base: 2, md: 20 }} py={{ base: 2, md: 4 }}>
            <Stack direction={{base: "column", md: "row"}} spacing={{base: 4}}>
                <VStack maxW={"200px"} spacing={{ base: 1, md: 2 }} align='stretch'>
                    <MenuItem label={"Absen"} redirect={PATH.CLASS + "qwerty" + "/attendance"} />
                    <MenuItem label={"Jadwal"} redirect={PATH.CLASS + "qwerty" + "/schedules"} />
                    <MenuItem label={"Siswa"} redirect={PATH.CLASS + "qwerty" + "/students"} />
                    <MenuItem label={"Orang Tua"} redirect={PATH.CLASS + "qwerty" + "/parents"} />
                    <MenuItem label={"Pengaturan"} redirect={PATH.CLASS + "qwerty" + "/setting"} />
                </VStack>
                <Box flex={1}>
                    <Card title={"Absen"}>
                        <Flex justifyContent={"space-between"}>
                            <Button size={{base: "xs", md: "sm"}}>Buat Absen</Button>
                            <Input size={{base: "xs", md: "sm"}} type="text" placeholder="25-06-2022" w={"auto"} />
                        </Flex>
                    </Card>
                </Box>
            </Stack>
        </Box>
    </Box>
}

export default Class;