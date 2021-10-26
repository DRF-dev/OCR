package file_content

func (f Model) FromBytesToString() string  {
	return string(f.Content)
}
