export const emailValidation = {
    required: "Must be filled",
    validate: (value: string) => {
        //eslint-disable-next-line
        if (!value.match(/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/)) {
            return "An email address is required."
        }
        return true
    }
}

export const passwordValidation = {
    required: "Must be filled",
    validate: (value: string) => {
        if (!value.match(/^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/)) {
            return "Password should contain at least eight characters, one letter and one number."
        }
        return true
    }
}

export const confirmPasswordValidation = (watchpass: string) => ({
    required: "Must be filled",
    validate: (value: string) => {
        if (watchpass !== value) {
            return "Password doesn't match"
        }
        return true
      }
    })
