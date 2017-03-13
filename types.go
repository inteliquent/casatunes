package casatunes

type RESTMediaItem struct {
  Flags int
  ID string
  GroupName string
  Title string
  Biography string
  Type int
  ArtworkURI string
  ArtworkRatio float32
  Album string
  Artists string
  Duration int64
  Genres string
  Released string
  Track string
  BitRate int
  ListenerCount int
  StationBand string
  StationDescription string
  StationFrequency string
  StationName string
  Rating int
  Description string
  ServiceName string
  ServiceLogoURI string
  TotalItems int
  SearchPlaceholderText string
  DisplayInfo []string
  ContextMenuItems []RESTContextMenuItem
}

type RESTContextMenuItem struct {
  Type int
  Title string
  Value string
}
