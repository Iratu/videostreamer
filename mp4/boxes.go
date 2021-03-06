package mp4

import (
	"bytes"
)

type Box interface {
	Name() string
	Data() []byte
	Children() []Box
}

type Sample struct {
	Duration uint32
	Size     uint32
	Scto     uint32
}

func BoxWrite(box Box) []byte {
	buf := &bytes.Buffer{}
	Write32(buf, 0)
	buf.Write([]byte(box.Name()))
	buf.Write(box.Data())
	for _, b := range box.Children() {
		buf.Write(BoxWrite(b))
	}
	data := buf.Bytes()
	WriteB32(data, uint32(len(data)))
	return data
}

type FtypBox struct {
	MajorBand       string
	MinorVersion    uint32
	CompatibleBands []string
}

func (box *FtypBox) Name() string {
	return "ftyp"
}

func (box *FtypBox) Children() []Box {
	return []Box{}
}

func (box *FtypBox) Data() []byte {
	buf := &bytes.Buffer{}
	buf.Write([]byte(box.MajorBand))
	Write32(buf, box.MinorVersion)
	for _, compat := range box.CompatibleBands {
		buf.Write([]byte(compat))
	}
	return buf.Bytes()
}

type MoovBox struct {
	BoxChildren []Box
}

func (box *MoovBox) Name() string {
	return "moov"
}

func (box *MoovBox) Children() []Box {
	return box.BoxChildren
}

func (box *MoovBox) Data() []byte {
	return []byte{}
}

type MvhdBox struct {
	CreationTime     uint64
	ModificationTime uint64
	Timescale        uint32
	Duration         uint64
	NextTrackId      uint32
}

func (box *MvhdBox) Name() string {
	return "mvhd"
}

func (box *MvhdBox) Children() []Box {
	return []Box{}
}

func (box *MvhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, uint32(box.CreationTime))
	Write32(buf, uint32(box.ModificationTime))
	Write32(buf, box.Timescale)
	Write32(buf, uint32(box.Duration))
	Write32(buf, 0x00010000)
	Write16(buf, 0x0100)
	Write16(buf, 0)

	Write32(buf, 0)
	Write32(buf, 0)

	Write32(buf, 0x00010000)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0x00010000)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0x40000000)

	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)

	Write32(buf, box.NextTrackId)
	return buf.Bytes()
}

type TrakBox struct {
	BoxChildren []Box
}

func (box *TrakBox) Name() string {
	return "trak"
}

func (box *TrakBox) Children() []Box {
	return box.BoxChildren
}

func (box *TrakBox) Data() []byte {
	return []byte{}
}

type TkhdBox struct {
	CreationTime     uint64
	ModificationTime uint64
	TrackId          uint32
	Duration         uint64
	Audio            bool
	Width            uint32
	Height           uint32
}

func (box *TkhdBox) Name() string {
	return "tkhd"
}

func (box *TkhdBox) Children() []Box {
	return []Box{}
}

func (box *TkhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0x3)
	Write32(buf, uint32(box.CreationTime))
	Write32(buf, uint32(box.ModificationTime))
	Write32(buf, box.TrackId)
	Write32(buf, 0)
	Write32(buf, uint32(box.Duration))
	Write32(buf, 0)
	Write32(buf, 0)

	Write16(buf, 0)
	Write16(buf, 0)
	if box.Audio {
		Write16(buf, 0x0100)
	} else {
		Write16(buf, 0)
	}
	Write16(buf, 0)

	Write32(buf, 0x00010000)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0x00010000)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0x40000000)

	Write32(buf, box.Width)
	Write32(buf, box.Height)

	return buf.Bytes()
}

type MdiaBox struct {
	BoxChildren []Box
}

func (box *MdiaBox) Name() string {
	return "mdia"
}

func (box *MdiaBox) Children() []Box {
	return box.BoxChildren
}

func (box *MdiaBox) Data() []byte {
	return []byte{}
}

type MdhdBox struct {
	CreationTime     uint64
	ModificationTime uint64
	Timescale        uint32
	Duration         uint64
}

func (box *MdhdBox) Name() string {
	return "mdhd"
}

func (box *MdhdBox) Children() []Box {
	return []Box{}
}

func (box *MdhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, uint32(box.CreationTime))
	Write32(buf, uint32(box.ModificationTime))
	Write32(buf, box.Timescale)
	Write32(buf, uint32(box.Duration))
	Write16(buf, 0)
	Write16(buf, 0)

	return buf.Bytes()
}

type HdlrBox struct {
	HandlerType string
	HandlerName string
}

func (box *HdlrBox) Name() string {
	return "hdlr"
}

func (box *HdlrBox) Children() []Box {
	return []Box{}
}

func (box *HdlrBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, 0)
	buf.Write([]byte(box.HandlerType))

	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)

	buf.Write([]byte(box.HandlerName))
	buf.Write([]byte{0})

	return buf.Bytes()
}

type MinfBox struct {
	BoxChildren []Box
}

func (box *MinfBox) Name() string {
	return "minf"
}

func (box *MinfBox) Children() []Box {
	return box.BoxChildren
}

func (box *MinfBox) Data() []byte {
	return []byte{}
}

type VmhdBox struct {
}

func (box *VmhdBox) Name() string {
	return "vmhd"
}

func (box *VmhdBox) Children() []Box {
	return []Box{}
}

func (box *VmhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 1)
	Write16(buf, 0)

	Write16(buf, 0)
	Write16(buf, 0)
	Write16(buf, 0)

	return buf.Bytes()
}

type SmhdBox struct {
}

func (box *SmhdBox) Name() string {
	return "smhd"
}

func (box *SmhdBox) Children() []Box {
	return []Box{}
}

func (box *SmhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 1)
	Write16(buf, 0)
	Write16(buf, 0)
	return buf.Bytes()
}

type DinfBox struct {
	BoxChildren []Box
}

func (box *DinfBox) Name() string {
	return "dinf"
}

func (box *DinfBox) Children() []Box {
	return box.BoxChildren
}

func (box *DinfBox) Data() []byte {
	return []byte{}
}

type DrefBox struct {
	BoxChildren []Box
}

func (box *DrefBox) Name() string {
	return "dref"
}

func (box *DrefBox) Children() []Box {
	return box.BoxChildren
}

func (box *DrefBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, uint32(len(box.BoxChildren)))
	return buf.Bytes()
}

type Url_Box struct {
}

func (box *Url_Box) Name() string {
	return "url "
}

func (box *Url_Box) Children() []Box {
	return []Box{}
}

func (box *Url_Box) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0x1)
	return buf.Bytes()
}

type StblBox struct {
	BoxChildren []Box
}

func (box *StblBox) Name() string {
	return "stbl"
}

func (box *StblBox) Children() []Box {
	return box.BoxChildren
}

func (box *StblBox) Data() []byte {
	return []byte{}
}

type StsdBox struct {
	BoxChildren []Box
}

func (box *StsdBox) Name() string {
	return "stsd"
}

func (box *StsdBox) Children() []Box {
	return box.BoxChildren
}

func (box *StsdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, uint32(len(box.BoxChildren)))

	return buf.Bytes()
}

type Avc1Box struct {
	DataReferenceIndex uint16
	Width              uint16
	Height             uint16
	Compressorname     string
	AvcC               *AvcCBox
}

func (box *Avc1Box) Name() string {
	return "avc1"
}

func (box *Avc1Box) Children() []Box {
	return []Box{box.AvcC}
}

func (box *Avc1Box) Data() []byte {
	buf := &bytes.Buffer{}

	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)

	Write16(buf, box.DataReferenceIndex)

	Write16(buf, 0)
	Write16(buf, 0)

	Write32(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)

	Write16(buf, box.Width)
	Write16(buf, box.Height)

	Write32(buf, 0x00480000)
	Write32(buf, 0x00480000)

	Write32(buf, 0)
	Write16(buf, 1)
	name := make([]byte, 32)
	copy(name, []byte(box.Compressorname))
	buf.Write(name)
	Write16(buf, 0x0018)
	Write16(buf, ^uint16(0))

	return buf.Bytes()
}

type AvcCBox struct {
	AvcCData []byte
}

func (box *AvcCBox) Name() string {
	return "avcC"
}

func (box *AvcCBox) Children() []Box {
	return []Box{}
}

func (box *AvcCBox) Data() []byte {
	return box.AvcCData
}

type Mp4aBox struct {
	DataReferenceIndex uint16
	SampleRate         uint32
	Esds               *EsdsBox
}

func (box *Mp4aBox) Name() string {
	return "mp4a"
}

func (box *Mp4aBox) Children() []Box {
	return []Box{box.Esds}
}

func (box *Mp4aBox) Data() []byte {
	buf := &bytes.Buffer{}

	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)
	Write8(buf, 0)

	Write16(buf, box.DataReferenceIndex)

	Write32(buf, 0)
	Write32(buf, 0)

	Write16(buf, 2)
	Write16(buf, 16)

	Write16(buf, 0)
	Write16(buf, 0)

	Write32(buf, box.SampleRate<<16)
	return buf.Bytes()
}

type EsdsBox struct {
	Bitrate   uint32
	Frequency uint32
}

func (box *EsdsBox) Name() string {
	return "esds"
}

func (box *EsdsBox) Children() []Box {
	return []Box{}
}

func (box *EsdsBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write8(buf, 0x03)
	Write8(buf, 0x19)
	Write16(buf, 0x00)
	Write8(buf, 0x00)
	Write8(buf, 0x04)
	Write8(buf, 0x11)
	Write8(buf, 0x40)
	Write8(buf, 0x15)
	Write8(buf, 0x00)
	Write8(buf, 0x06)
	Write8(buf, 0x00)
	Write32(buf, box.Bitrate)
	Write32(buf, box.Bitrate)
	Write8(buf, 0x05)
	Write8(buf, 0x02)
	Write8(buf, 0x12)
	Write8(buf, 0x10)
	Write8(buf, 0x06)
	Write8(buf, 0x01)
	Write8(buf, 0x02)

	return buf.Bytes()
}

type SttsBox struct {
}

func (box *SttsBox) Name() string {
	return "stts"
}

func (box *SttsBox) Children() []Box {
	return []Box{}
}

func (box *SttsBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, 0)
	return buf.Bytes()
}

type StscBox struct {
}

func (box *StscBox) Name() string {
	return "stsc"
}

func (box *StscBox) Children() []Box {
	return []Box{}
}

func (box *StscBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, 0)
	return buf.Bytes()
}

type StszBox struct {
}

func (box *StszBox) Name() string {
	return "stsz"
}

func (box *StszBox) Children() []Box {
	return []Box{}
}

func (box *StszBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, 0)
	Write32(buf, 0)
	return buf.Bytes()
}

type StcoBox struct {
}

func (box *StcoBox) Name() string {
	return "stco"
}

func (box *StcoBox) Children() []Box {
	return []Box{}
}

func (box *StcoBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, 0)
	return buf.Bytes()
}

type SidxBox struct {
	ReferenceId        uint32
	Timescale          uint32
	PresentationTime   uint64
	ReferenceSize      uint32
	SubsegmentDuration uint32
	Keyframe           bool
}

func (box *SidxBox) Name() string {
	return "sidx"
}

func (box *SidxBox) Children() []Box {
	return []Box{}
}

func (box *SidxBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 1)
	Write24(buf, 0)
	Write32(buf, box.ReferenceId)
	Write32(buf, box.Timescale)
	Write64(buf, box.PresentationTime)
	Write64(buf, 0)
	Write16(buf, 0)
	Write16(buf, 1)
	Write32(buf, box.ReferenceSize&0x7fffffff)
	Write32(buf, box.SubsegmentDuration)
	if box.Keyframe {
		Write32(buf, 1<<31)
	} else {
		Write32(buf, 0)
	}
	return buf.Bytes()
}

type MoofBox struct {
	BoxChildren []Box
}

func (box *MoofBox) Name() string {
	return "moof"
}

func (box *MoofBox) Children() []Box {
	return box.BoxChildren
}

func (box *MoofBox) Data() []byte {
	return []byte{}
}

type MfhdBox struct {
	Sequence uint32
}

func (box *MfhdBox) Name() string {
	return "mfhd"
}

func (box *MfhdBox) Children() []Box {
	return []Box{}
}

func (box *MfhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, box.Sequence)
	return buf.Bytes()
}

type TrafBox struct {
	BoxChildren []Box
}

func (box *TrafBox) Name() string {
	return "traf"
}

func (box *TrafBox) Children() []Box {
	return box.BoxChildren
}

func (box *TrafBox) Data() []byte {
	return []byte{}
}

type TfhdBox struct {
	TrackId uint32
	Flags   uint32
}

func (box *TfhdBox) Name() string {
	return "tfhd"
}

func (box *TfhdBox) Children() []Box {
	return []Box{}
}

func (box *TfhdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0x20000|0x20)
	Write32(buf, box.TrackId)
	Write32(buf, box.Flags)
	return buf.Bytes()
}

type TfdtBox struct {
	BaseMediaDecodeTime uint64
}

func (box *TfdtBox) Name() string {
	return "tfdt"
}

func (box *TfdtBox) Children() []Box {
	return []Box{}
}

func (box *TfdtBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 1)
	Write24(buf, 0)
	Write64(buf, box.BaseMediaDecodeTime)
	return buf.Bytes()
}

type TrunBox struct {
	Samples []*Sample
}

func (box *TrunBox) Name() string {
	return "trun"
}

func (box *TrunBox) Children() []Box {
	return []Box{}
}

func (box *TrunBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0x800|0x200|0x100|0x1)
	Write32(buf, uint32(len(box.Samples)))
	Write32(buf, 0)
	for _, s := range box.Samples {
		Write32(buf, s.Duration)
		Write32(buf, s.Size)
		Write32(buf, s.Scto)
	}
	return buf.Bytes()
}

type MdatBox struct {
	BoxData []byte
}

func (box *MdatBox) Name() string {
	return "mdat"
}

func (box *MdatBox) Children() []Box {
	return []Box{}
}

func (box *MdatBox) Data() []byte {
	return box.BoxData
}

type MvexBox struct {
	BoxChildren []Box
}

func (box *MvexBox) Name() string {
	return "mvex"
}

func (box *MvexBox) Children() []Box {
	return box.BoxChildren
}

func (box *MvexBox) Data() []byte {
	return []byte{}
}

type MehdBox struct {
	TimeScale uint32
}

func (box *MehdBox) Name() string {
	return "mehd"
}

func (box *MehdBox) Children() []Box {
	return []Box{}
}

func (box *MehdBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, box.TimeScale)
	return buf.Bytes()
}

type TrexBox struct {
	TrackId uint32
	Defdur  uint32
}

func (box *TrexBox) Name() string {
	return "trex"
}

func (box *TrexBox) Children() []Box {
	return []Box{}
}

func (box *TrexBox) Data() []byte {
	buf := &bytes.Buffer{}
	Write8(buf, 0)
	Write24(buf, 0)
	Write32(buf, box.TrackId)
	Write32(buf, box.Defdur)
	Write32(buf, 1)
	Write32(buf, 0)
	Write32(buf, 0)
	return buf.Bytes()
}
