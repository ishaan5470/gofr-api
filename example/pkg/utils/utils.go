package utils

func ParseBody(c gofr.Context, x interface{}) error {
    if err := c.BodyParser(x); err != nil {
        // Handle parsing error based on your needs
        return err
    }
    return nil
}
