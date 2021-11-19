package asig

type Material struct {

	/** List of all material properties loaded. */
	Properties []*MaterialProperty

	/** Storage allocated */
	AllocatedStorage uint
}

type MaterialProperty struct {

	//Specifies the name of the property (aka key). Keys are generally case insensitive.
	name string

	/** Textures: Specifies their exact usage semantic.
	 * For non-texture properties, this member is always 0 (aka TextureTypeNone).
	 */
	Semantic TextureType

	/** Textures: Specifies the index of the texture.
	 *  For non-texture properties, this member is always 0.
	 */
	Index uint

	/** Type information for the property.
	 *
	 * Defines the data layout inside the data buffer. This is used
	 * by the library internally to perform debug checks and to
	 * utilize proper type conversions.
	 * (It's probably a hacky solution, but it works.)
	 */
	TypeInfo MatPropertyTypeInfo

	/** Binary buffer to hold the property's value.
	 * The size of the buffer is always mDataLength.
	 */
	Data []byte
}
