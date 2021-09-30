import React from 'react';
import {Formik, Form, Field} from 'formik'
import {
  ChakraProvider,
  Box,
  Button,
  Input,
  Grid,
  FormControl,
  InputGroup,
  InputRightElement,
  Center
} from '@chakra-ui/react';

import { createStandaloneToast } from "@chakra-ui/react"
import { extendTheme } from "@chakra-ui/react"
import { ColorModeSwitcher } from './ColorModeSwitcher';
import { mode } from "@chakra-ui/theme-tools"

const extended_theme = extendTheme({
  styles: {
    global: (props) => ({
      body: {
        bg: mode("linear-gradient(#7928CA, #FF0080) fixed", "linear-gradient(147deg, #000000 0%, #04317f 94%) fixed")(props)
      },
    }),
  },
})

function showCreationToast(toast, newUrl){
  toast({
    title: "Short link created.",
    description: "You can access original destination using: " + newUrl,
    status: "info",
    variant: "solid",
    duration: 10000,
    isClosable: true,
  })
}

function showFailureToast(toast){
  toast({
    title: "Operation failed.",
    description: "Could not create short link.",
    status: "error",
    variant: "solid",
    duration: 10000,
    isClosable: true,
  })
}

function App() {
  const toast = createStandaloneToast(extended_theme)
  return (
    <ChakraProvider theme={extended_theme}>
      <Box textAlign="center" fontSize="xl">
        <Grid>
          <ColorModeSwitcher justifySelf="flex-end" />
          <Center minH="80vh">
            <Formik
              initialValues={{ url: "" }}
              onSubmit={(values, actions) => {

                  const apiUrl = process.env.REACT_APP_API_URL

                  fetch(apiUrl + '/url', {
                    method: 'post',
                    headers: {
                      'Accept': 'application/json, text/plain, */*',
                      'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({original_url: values.url})
                  })
                  .then(response => response.json())
                  .then(data => {
                    if (!data.new_url) {
                      showFailureToast(toast)
                      return
                    }
                    showCreationToast(toast, data.new_url)
                  })
                  .catch((error) => {
                    showFailureToast(toast)
                  })
                  .finally(function(){
                    actions.setSubmitting(false)
                  });
              }}
            >
              {({ isSubmitting }) => (
                <Form>
                  <Field name="url">
                    {({ field, form }) => (
                    <FormControl width="xl">
                      <InputGroup>
                        <Input
                        {...field} id="url" textColor="white" placeholder="Enter the link and shorten it"/>
                        <InputRightElement width="6.5rem">
                          <Button type="submit" isLoading={isSubmitting} h="80%" opacity="70%" margin="0.5rem">
                            Shorten
                          </Button>
                        </InputRightElement>
                      </InputGroup>
                    </FormControl>
                    )}
                  </Field>
                </Form>
              )}
            </Formik>
          </Center>
        </Grid>
      </Box>
    </ChakraProvider>
  );
}

export default App;
